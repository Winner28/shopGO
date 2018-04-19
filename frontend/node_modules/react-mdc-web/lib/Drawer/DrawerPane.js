'use strict';

Object.defineProperty(exports, "__esModule", {
  value: true
});

var _createClass = function () { function defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ("value" in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } } return function (Constructor, protoProps, staticProps) { if (protoProps) defineProperties(Constructor.prototype, protoProps); if (staticProps) defineProperties(Constructor, staticProps); return Constructor; }; }();

var _propTypes = require('prop-types');

var _propTypes2 = _interopRequireDefault(_propTypes);

var _react = require('react');

var _react2 = _interopRequireDefault(_react);

var _classnames = require('classnames');

var _classnames2 = _interopRequireDefault(_classnames);

var _constants = require('./constants');

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }

function _possibleConstructorReturn(self, call) { if (!self) { throw new ReferenceError("this hasn't been initialised - super() hasn't been called"); } return call && (typeof call === "object" || typeof call === "function") ? call : self; }

function _inherits(subClass, superClass) { if (typeof superClass !== "function" && superClass !== null) { throw new TypeError("Super expression must either be null or a function, not " + typeof superClass); } subClass.prototype = Object.create(superClass && superClass.prototype, { constructor: { value: subClass, enumerable: false, writable: true, configurable: true } }); if (superClass) Object.setPrototypeOf ? Object.setPrototypeOf(subClass, superClass) : subClass.__proto__ = superClass; }

var DrawerPane = function (_React$PureComponent) {
  _inherits(DrawerPane, _React$PureComponent);

  function DrawerPane(props) {
    _classCallCheck(this, DrawerPane);

    var _this = _possibleConstructorReturn(this, (DrawerPane.__proto__ || Object.getPrototypeOf(DrawerPane)).call(this, props));

    _this.handleTransitionend = _this.handleTransitionend.bind(_this);
    _this.handleTouchstart = _this.handleTouchstart.bind(_this);
    return _this;
  }

  _createClass(DrawerPane, [{
    key: 'componentWillReceiveProps',
    value: function componentWillReceiveProps(next) {
      var _props = this.props,
          animating = _props.animating,
          position = _props.position;

      if (animating !== next.animating && next.animating) {
        this.native.addEventListener('transitionend', this.handleTransitionend);
      }

      if (position !== next.position) {
        this.setTransformPosition(next.position);
      }
    }
  }, {
    key: 'getTransformPropertyName',
    value: function getTransformPropertyName() {
      if (!this.transformPropertyName) {
        var el = document.createElement('div');
        this.transformPropertyName = 'transform' in el.style ? 'transform' : '-webkit-transform';
      }
      return this.transformPropertyName;
    }
  }, {
    key: 'setTransformPosition',
    value: function setTransformPosition(position) {
      var propertyName = this.getTransformPropertyName();
      var transform = position === null ? null : 'translateX(' + position + 'px)';
      this.native.style.setProperty(propertyName, transform);
    }
  }, {
    key: 'handleTransitionend',
    value: function handleTransitionend() {
      var onTransitionend = this.props.onTransitionend;

      this.native.removeEventListener('transitionend', this.handleTransitionend);
      onTransitionend();
    }
  }, {
    key: 'handleTouchstart',
    value: function handleTouchstart(event) {
      var onTouchstart = this.props.onTouchstart;

      var width = this.native.offsetWidth;
      onTouchstart(event, width);
    }
  }, {
    key: 'render',
    value: function render() {
      var _this2 = this;

      var _props2 = this.props,
          className = _props2.className,
          children = _props2.children,
          onClick = _props2.onClick;


      return _react2.default.createElement(
        'nav',
        {
          className: (0, _classnames2.default)(_constants.ROOT + '__drawer', className),
          ref: function ref(native) {
            _this2.native = native;
          },
          onClick: onClick,
          onTouchStart: this.handleTouchstart
        },
        children
      );
    }
  }]);

  return DrawerPane;
}(_react2.default.PureComponent);

DrawerPane.propTypes = {
  children: _propTypes2.default.node,
  className: _propTypes2.default.string,
  onClick: _propTypes2.default.func,
  onTouchstart: _propTypes2.default.func,
  onTransitionend: _propTypes2.default.func,
  animating: _propTypes2.default.bool,
  position: _propTypes2.default.number
};
exports.default = DrawerPane;