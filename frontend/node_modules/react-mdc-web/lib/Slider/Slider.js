'use strict';

Object.defineProperty(exports, "__esModule", {
  value: true
});

var _extends = Object.assign || function (target) { for (var i = 1; i < arguments.length; i++) { var source = arguments[i]; for (var key in source) { if (Object.prototype.hasOwnProperty.call(source, key)) { target[key] = source[key]; } } } return target; };

var _createClass = function () { function defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ("value" in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } } return function (Constructor, protoProps, staticProps) { if (protoProps) defineProperties(Constructor.prototype, protoProps); if (staticProps) defineProperties(Constructor, staticProps); return Constructor; }; }();

var _propTypes = require('prop-types');

var _propTypes2 = _interopRequireDefault(_propTypes);

var _react = require('react');

var _react2 = _interopRequireDefault(_react);

var _classnames2 = require('classnames');

var _classnames3 = _interopRequireDefault(_classnames2);

var _mdc = require('@material/slider/dist/mdc.slider');

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

function _defineProperty(obj, key, value) { if (key in obj) { Object.defineProperty(obj, key, { value: value, enumerable: true, configurable: true, writable: true }); } else { obj[key] = value; } return obj; }

function _objectWithoutProperties(obj, keys) { var target = {}; for (var i in obj) { if (keys.indexOf(i) >= 0) continue; if (!Object.prototype.hasOwnProperty.call(obj, i)) continue; target[i] = obj[i]; } return target; }

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }

function _possibleConstructorReturn(self, call) { if (!self) { throw new ReferenceError("this hasn't been initialised - super() hasn't been called"); } return call && (typeof call === "object" || typeof call === "function") ? call : self; }

function _inherits(subClass, superClass) { if (typeof superClass !== "function" && superClass !== null) { throw new TypeError("Super expression must either be null or a function, not " + typeof superClass); } subClass.prototype = Object.create(superClass && superClass.prototype, { constructor: { value: subClass, enumerable: false, writable: true, configurable: true } }); if (superClass) Object.setPrototypeOf ? Object.setPrototypeOf(subClass, superClass) : subClass.__proto__ = superClass; }

var DISCRETE = 'mdc-slider--discrete';
var SHOW_MARKERS = 'mdc-slider--display-markers';

var Slider = function (_Component) {
  _inherits(Slider, _Component);

  function Slider() {
    _classCallCheck(this, Slider);

    return _possibleConstructorReturn(this, (Slider.__proto__ || Object.getPrototypeOf(Slider)).apply(this, arguments));
  }

  _createClass(Slider, [{
    key: 'componentDidMount',
    value: function componentDidMount() {
      var _this2 = this;

      this.slider = new _mdc.MDCSlider(this.node);
      this.setup(this.props);

      this.slider.listen('MDCSlider:change', function () {
        return _this2.handleMDCSliderChange();
      });
      this.slider.listen('MDCSlider:input', function () {
        return _this2.handleMDCSliderInput();
      });
    }
  }, {
    key: 'componentWillReceiveProps',
    value: function componentWillReceiveProps(newProps) {
      this.setup(newProps);
    }
  }, {
    key: 'componentWillUnmount',
    value: function componentWillUnmount() {
      var _this3 = this;

      this.slider.unlisten('MDCSlider:change', function () {
        return _this3.handleMDCSliderChange();
      });
      this.slider.unlisten('MDCSlider:input', function () {
        return _this3.handleMDCSliderInput();
      });
      this.slider.destroy();
    }
  }, {
    key: 'setup',
    value: function setup(props) {
      this.slider.disabled = props.disabled;
      if (props.min) {
        this.slider.min = props.min;
      }
      if (props.max) {
        this.slider.max = props.max;
      }
      if (props.step) {
        this.slider.step = props.step;
      }
      this.slider.value = props.value;
    }
  }, {
    key: 'handleMDCSliderInput',
    value: function handleMDCSliderInput() {
      if (this.props.onInput) {
        this.props.onInput(this.slider.value);
      }
    }
  }, {
    key: 'handleMDCSliderChange',
    value: function handleMDCSliderChange() {
      if (this.props.onChange) {
        this.props.onChange(this.slider.value);
      }
    }
  }, {
    key: 'render',
    value: function render() {
      var _this4 = this,
          _classnames;

      var _props = this.props,
          className = _props.className,
          value = _props.value,
          min = _props.min,
          max = _props.max,
          step = _props.step,
          discrete = _props.discrete,
          _props$showMarkers = _props.showMarkers,
          showMarkers = _props$showMarkers === undefined ? false : _props$showMarkers,
          otherProps = _objectWithoutProperties(_props, ['className', 'value', 'min', 'max', 'step', 'discrete', 'showMarkers']);

      return _react2.default.createElement(
        'div',
        _extends({
          ref: function ref(node) {
            _this4.node = node;
          },
          className: (0, _classnames3.default)('mdc-slider', (_classnames = {}, _defineProperty(_classnames, SHOW_MARKERS, showMarkers), _defineProperty(_classnames, DISCRETE, discrete), _classnames), className)
        }, otherProps),
        _react2.default.createElement(
          'div',
          { className: 'mdc-slider__track-container' },
          _react2.default.createElement('div', { className: 'mdc-slider__track' }),
          discrete && showMarkers && _react2.default.createElement('div', { className: 'mdc-slider__track-marker-container' })
        ),
        _react2.default.createElement(
          'div',
          { className: 'mdc-slider__thumb-container' },
          discrete && _react2.default.createElement(
            'div',
            { className: 'mdc-slider__pin' },
            _react2.default.createElement('span', { className: 'mdc-slider__pin-value-marker' })
          ),
          _react2.default.createElement(
            'svg',
            { className: 'mdc-slider__thumb', width: '21', height: '21' },
            _react2.default.createElement('circle', { cx: '10.5', cy: '10.5', r: '7.875' })
          ),
          _react2.default.createElement('div', { className: 'mdc-slider__focus-ring' })
        )
      );
    }
  }]);

  return Slider;
}(_react.Component);

Slider.propTypes = {
  className: _propTypes2.default.string,
  value: _propTypes2.default.number.isRequired,
  min: _propTypes2.default.number,
  max: _propTypes2.default.number,
  step: _propTypes2.default.number,
  // Per MDC documentation providing step doesn't imply discrete, so an explicit property
  discrete: _propTypes2.default.bool,
  showMarkers: _propTypes2.default.bool,
  onInput: _propTypes2.default.func,
  onChange: _propTypes2.default.func
};
exports.default = Slider;