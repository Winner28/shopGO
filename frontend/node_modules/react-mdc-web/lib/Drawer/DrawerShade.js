'use strict';

Object.defineProperty(exports, "__esModule", {
  value: true
});

var _createClass = function () { function defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ("value" in descriptor) descriptor.writable = true; Object.defineProperty(target, descriptor.key, descriptor); } } return function (Constructor, protoProps, staticProps) { if (protoProps) defineProperties(Constructor.prototype, protoProps); if (staticProps) defineProperties(Constructor, staticProps); return Constructor; }; }();

var _propTypes = require('prop-types');

var _propTypes2 = _interopRequireDefault(_propTypes);

var _react = require('react');

var _react2 = _interopRequireDefault(_react);

var _classnames2 = require('classnames');

var _classnames3 = _interopRequireDefault(_classnames2);

var _constants = require('./constants');

var _isSupportCssCustomProperties = require('../utils/isSupportCssCustomProperties');

var _isSupportCssCustomProperties2 = _interopRequireDefault(_isSupportCssCustomProperties);

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

function _defineProperty(obj, key, value) { if (key in obj) { Object.defineProperty(obj, key, { value: value, enumerable: true, configurable: true, writable: true }); } else { obj[key] = value; } return obj; }

function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }

function _possibleConstructorReturn(self, call) { if (!self) { throw new ReferenceError("this hasn't been initialised - super() hasn't been called"); } return call && (typeof call === "object" || typeof call === "function") ? call : self; }

function _inherits(subClass, superClass) { if (typeof superClass !== "function" && superClass !== null) { throw new TypeError("Super expression must either be null or a function, not " + typeof superClass); } subClass.prototype = Object.create(superClass && superClass.prototype, { constructor: { value: subClass, enumerable: false, writable: true, configurable: true } }); if (superClass) Object.setPrototypeOf ? Object.setPrototypeOf(subClass, superClass) : subClass.__proto__ = superClass; }

var DrawerShade = function (_React$PureComponent) {
  _inherits(DrawerShade, _React$PureComponent);

  function DrawerShade() {
    _classCallCheck(this, DrawerShade);

    return _possibleConstructorReturn(this, (DrawerShade.__proto__ || Object.getPrototypeOf(DrawerShade)).apply(this, arguments));
  }

  _createClass(DrawerShade, [{
    key: 'componentWillReceiveProps',
    value: function componentWillReceiveProps(next) {
      var opacity = this.props.opacity;

      if (opacity !== next.opacity && (0, _isSupportCssCustomProperties2.default)()) {
        this.native.style.setProperty('--' + _constants.ROOT + '-opacity', next.opacity);
      }
    }
  }, {
    key: 'render',
    value: function render() {
      var _classnames,
          _this2 = this;

      var _props = this.props,
          animating = _props.animating,
          className = _props.className,
          children = _props.children,
          onClick = _props.onClick,
          onTouchend = _props.onTouchend,
          onTouchmove = _props.onTouchmove,
          open = _props.open;


      return _react2.default.createElement(
        'aside',
        {
          className: (0, _classnames3.default)(_constants.ROOT, (_classnames = {}, _defineProperty(_classnames, _constants.ROOT + '--open', open), _defineProperty(_classnames, _constants.ROOT + '--animating', animating), _classnames), className),
          ref: function ref(native) {
            _this2.native = native;
          },
          onClick: onClick,
          onTouchEnd: onTouchend,
          onTouchMove: onTouchmove
        },
        children
      );
    }
  }]);

  return DrawerShade;
}(_react2.default.PureComponent);

DrawerShade.propTypes = {
  animating: _propTypes2.default.bool,
  children: _propTypes2.default.node,
  className: _propTypes2.default.string,
  open: _propTypes2.default.bool,
  onClick: _propTypes2.default.func,
  onTouchmove: _propTypes2.default.func,
  onTouchend: _propTypes2.default.func,
  opacity: _propTypes2.default.number
};
exports.default = DrawerShade;